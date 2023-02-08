# Load the memory dump into a byte array
$memoryDump = Get-Content -Path <path_to_memory_dump_file> -Encoding Byte

# Set the starting offset for the search
$offset = 0

# Repeat the search until all PE files have been extracted
while ($offset -lt $memoryDump.Length) {

  # Check if the current offset is the start of a PE file
  if ($memoryDump[$offset..$offset+2] -eq [Byte[]](0x4d, 0x5a, 0x90)) {

    # Extract the PE header to find the size of the file
    $peHeader = $memoryDump[$offset..$offset+23]
    $peSize = [BitConverter]::ToInt32($peHeader[0x14..0x17], 0)

    # Extract the PE file
    $peFile = $memoryDump[$offset..$offset+$peSize-1]

    # Save the PE file to disk
    $peFile | Set-Content -Path "PE_File_$offset.exe" -Encoding Byte

    # Move to the next offset
    $offset = $offset + $peSize
  } else {
    $offset = $offset + 1
  }
}

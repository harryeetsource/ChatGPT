Get-Process | ForEach-Object {
    $process = $_
    $processName = $process.ProcessName
    $pid = $process.Id

    # create a dump file with the process name and PID
    $dumpFile = "$processName-$pid.dmp"

    # create a memory dump of the process
    [Diagnostics.Process]::GetCurrentProcess().Refresh()
    [Diagnostics.Process]::GetProcessById($pid).Refresh()
    [Diagnostics.Process]::GetProcessById($pid).Dump("$dumpFile")

    # check if the dump file was created
    if (Test-Path $dumpFile) {
        Write-Output "Memory dump of $processName (PID: $pid) created successfully: $dumpFile"
    } else {
        Write-Output "Failed to create memory dump of $processName (PID: $pid)"
    }
}

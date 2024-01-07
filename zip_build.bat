@echo off
setlocal

:: Extract the productVersion from wails.json using PowerShell
for /f "delims=" %%i in ('powershell -Command "Get-Content wails.json | ConvertFrom-Json | Select -ExpandProperty Info | Select -ExpandProperty productVersion"') do set productVersion=%%i

:: Display the extracted version
echo Product Version: %productVersion%

set filename=mmd_rover_%productVersion%.exe

@REM wails build -trimpath -clean -o %filename%
wails build -trimpath -o %filename%

endlocal

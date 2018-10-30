@echo off

if "%BUILD_TYPE%"=="release" (
    adb install -r %ROOT_DIR%\generated\apks\app-release.apk
) else (
    adb install -r %ROOT_DIR%\generated\apks\app-debug.apk
)
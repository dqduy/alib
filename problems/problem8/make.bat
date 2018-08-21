@echo off

echo ===========================================================
echo                    BUILD APPLICATION
echo ===========================================================

call config.bat

echo Load build tools
echo NDK: %NDK_PATH%
echo SDK: %SDK_PATH%

echo Load current project
echo Project Root:          %ROOT_DIR%
echo Native directory:      %NATIVE_DIR%
echo Java directory:        %JAVA_DIR%
echo Script directory:      %SCRIPT_DIR%
echo Generated directory:   %GEN_DIR%

rem Processor Architecture, and build type
set P_ARCH=%~1
set BUILD_TYPE=%~2
set INSTALL=%~3

echo ===========================================================
echo                    MAKE NATIVE
echo ===========================================================

rem call %SCRIPT_DIR%\make_native.bat %P_ARCH% %BUILD_TYPE%

call %SCRIPT_DIR%\postProcessing.bat

echo ===========================================================
echo                    MAKE APK
echo ===========================================================

call %SCRIPT_DIR%\make_java.bat %P_ARCH% %BUILD_TYPE%


if "%~2"=="release" (
     xcopy /y %ROOT_DIR%generated\_asproject\app\build\outputs\apk\app-release.apk %ROOT_DIR%\generated\apks\
) else (
      xcopy /y %ROOT_DIR%generated\_asproject\app\build\outputs\apk\app-debug.apk %ROOT_DIR%\generated\apks\
)  

echo ===========================================================
echo                    INSTALL APK
echo ===========================================================

if "%~2"=="release" (
    adb install -r %ROOT_DIR%\generated\apks\app-release.apk
) else (
    adb install -r %ROOT_DIR%\generated\apks\app-debug.apk
)
 
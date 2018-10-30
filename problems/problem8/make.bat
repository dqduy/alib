@echo off

echo ===========================================================
echo                    BUILD APPLICATION
echo ===========================================================

call config.bat

rem Processor Architecture, and build type
set P_ARCH=%~1
set BUILD_TYPE=%~2
set INSTALL=%~3

echo Load build tools
echo NDK: %NDK_PATH%
echo SDK: %SDK_PATH%

echo Load current project
echo Project Root:          %ROOT_DIR%
echo Native directory:      %NATIVE_DIR%
echo Java directory:        %JAVA_DIR%
echo Script directory:      %SCRIPT_DIR%
echo Generated directory:   %GEN_DIR%

echo ARCH:          %P_ARCH%
echo BUILD_TYPE:    %BUILD_TYPE%

echo ===========================================================
echo                    MAKE NATIVE
echo ===========================================================

rem call %SCRIPT_DIR%\make_native.bat %P_ARCH% %BUILD_TYPE%

call %SCRIPT_DIR%\postProcessing.bat

echo ===========================================================
echo                    MAKE APK
echo ===========================================================

call %SCRIPT_DIR%\make_java.bat %P_ARCH% %BUILD_TYPE%



if "%BUILD_TYPE%"=="release" (
     xcopy /y %ROOT_DIR%generated\_asproject\app\build\outputs\apk\app-release.apk %ROOT_DIR%\generated\apks\
) else (
      xcopy /y %ROOT_DIR%generated\_asproject\app\build\outputs\apk\app-debug.apk %ROOT_DIR%\generated\apks\
)  

if "%INSTALL%"==1 (
    echo ===========================================================
    echo                    INSTALL APK
    echo ===========================================================

    call install.bat
)
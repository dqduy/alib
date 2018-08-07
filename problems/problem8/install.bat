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
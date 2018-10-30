if "%~2"=="release" (
    call %NDK_PATH%\ndk-build.cmd -C %NATIVE_DIR%\jni
    rem echo "if"
) else (
    rem echo "else"
    call %NDK_PATH%\ndk-build.cmd -C %NATIVE_DIR%\jni
)    



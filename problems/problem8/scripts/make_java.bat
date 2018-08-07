if "%~2"=="release" (
    call %JAVA_DIR%\project\gradlew.bat assembleRelease -p %ROOT_DIR%\generated\_asproject 
    
) else (
    call %JAVA_DIR%\project\gradlew.bat assembleDebug -p %ROOT_DIR%\generated\_asproject
)    




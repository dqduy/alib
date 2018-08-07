rem project Root

rem Copy libhello.so to java project
xcopy /s /e /y /i %NATIVE_DIR%\libs %JAVA_DIR%\project\app\src\main\libs

rem Copy java project to generated dir

xcopy /s /e /y /i %JAVA_DIR%\project %ROOT_DIR%\generated\_asproject
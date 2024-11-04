@echo off
SET SERVICE_NAME=MyGoService
SET DISPLAY_NAME=My Go Windows Service
SET SERVICE_PATH=%~dp0myservice.exe

IF "%1"=="install" (
    sc create "%SERVICE_NAME%" binPath= "%SERVICE_PATH%" DisplayName= "%DISPLAY_NAME%" start= auto
    sc description "%SERVICE_NAME%" "My Go Windows Service Description"
    echo Сервис "%SERVICE_NAME%" успешно установлен.
    GOTO :EOF
)

IF "%1"=="start" (
    sc start "%SERVICE_NAME%"
    GOTO :EOF
)

IF "%1"=="stop" (
    sc stop "%SERVICE_NAME%"
    GOTO :EOF
)

IF "%1"=="status" (
    sc query "%SERVICE_NAME%"
    GOTO :EOF
)

IF "%1"=="remove" (
    sc stop "%SERVICE_NAME%"
    sc delete "%SERVICE_NAME%"
    echo Сервис "%SERVICE_NAME%" успешно удален.
    GOTO :EOF
)

:USAGE
echo Использование:
echo    service.bat install   - установить сервис
echo    service.bat start     - запустить сервис
echo    service.bat stop      - остановить сервис
echo    service.bat status    - проверить статус сервиса
echo    service.bat remove    - удалить сервис
GOTO :EOF

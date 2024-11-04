@echo off
REM Установка переменных
SET SERVICE_NAME=MyGoService
SET BUILD_PATH=%~dp0
SET DEPLOY_PATH=C:\ProgramData\MyService
SET EXECUTABLE_NAME=myservice.exe

REM Проверка наличия Go в PATH
go version >nul 2>&1
IF ERRORLEVEL 1 (
    echo Go не найден в PATH. Убедитесь, что Go установлен и добавлен в переменную среды PATH.
    GOTO :EOF
)

REM Сборка проекта Go
echo Сборка Go-проекта...
go build -o "%BUILD_PATH%%EXECUTABLE_NAME%" ./cmd/service
IF ERRORLEVEL 1 (
    echo Сборка не удалась.
    GOTO :EOF
)
echo Сборка успешно завершена.

REM Создание директории для развертывания, если она не существует
IF NOT EXIST "%DEPLOY_PATH%" (
    echo Создание директории для развертывания "%DEPLOY_PATH%"
    mkdir "%DEPLOY_PATH%"
)

REM Копирование исполняемого файла
echo Копирование исполняемого файла в "%DEPLOY_PATH%"
copy /Y "%BUILD_PATH%%EXECUTABLE_NAME%" "%DEPLOY_PATH%"

REM Копирование config.ini
echo Копирование config.ini в "%DEPLOY_PATH%"
copy /Y "%BUILD_PATH%config.ini" "%DEPLOY_PATH%"

REM Копирование service.bat (если необходимо)
echo Копирование service.bat в "%DEPLOY_PATH%"
copy /Y "%BUILD_PATH%service.bat" "%DEPLOY_PATH%"

REM Развертывание завершено
echo Развертывание завершено. Файлы сервиса скопированы в "%DEPLOY_PATH%".
GOTO :EOF

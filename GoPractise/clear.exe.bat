@echo off 

for /R %%s in (*.exe) do (
	echo %%s
	del %%s
)

REM pause

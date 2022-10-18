@echo off
title kill proc by port
:start1
set duankou=%1
setlocal enabledelayedexpansion
for /f "delims=  tokens=1" %%a in ('netstat -aon ^| findstr "%duankou%" ^|findstr "LISTENING"') do (
	set a1=%%a
	set isCanUse=0
	for /f "tokens=2 delims= " %%b in ("!a1!") do (
		set a2=%%b
		for /f "tokens=2 delims=:" %%c in ("!a2!") do (
			set a3=%%c
			if %duankou%==!a3! (
				set isCanUse=1
			)
		)
		if !isCanUse!==1 (
			echo 查询到匹配的进程信息--!a1!
			for /f "tokens=5 delims= " %%d in ("!a1!") do (
				set killPid=%%d
				for /f "tokens=1 delims= " %%e in ('tasklist^|findstr  !killPid!') do (
					taskkill /f /t /im  %%e
				)
			)
		)
	)
)
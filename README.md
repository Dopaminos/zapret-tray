# zapret-tray
простая утилита для включения/выключения [zapret от Snowy-Fluffy](https://github.com/Snowy-Fluffy/zapret.installer) одной кнопкой в системном трее.

из-за моей конфигурации VPN периодически приходится отключать zapret, чтобы получить доступ к некоторым ресурсам.  
эта программка решает проблему с помощью включения-выключения zapret-а одним кликом по иконке в трее.
### build

```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o zapret-tray
```
### credits
https://github.com/energye/systray
https://github.com/Snowy-Fluffy/zapret.installer

`прогу я делал чисто для себя, поэтому её работа тестировалась только на арче`
## 自动更新壁纸

### 思路
1. 每周一读取指定文件夹，然后更改名称(1-7)，达到每天更换壁纸的效果
2. 根据名称读取文件的全路径
3. 利用win32api设置桌面背景

### 发布
pyinstall ./autoSwitchDeskPic.py

### 使用
1. 进入 C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup  
2. 把执行文件放进去
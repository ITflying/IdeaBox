"""
能够每天自动切换桌面背景
"""
import win32gui, win32api, win32con
import os
import random
from datetime import datetime


# 设置桌面背景
def set_wallpaper(imagePath):
    key = win32api.RegOpenKey(win32con.HKEY_CURRENT_USER, "Control Panel\\Desktop", 0, win32con.KEY_SET_VALUE)
    win32api.RegSetValueEx(key, "WallpaperStyle", 0, win32con.REG_SZ, "10")  # 修改背景风格：拉伸在注册表中的数值是2，适应是6，填充是10，而平铺和居中都是0
    win32api.RegSetValueEx(key, "TileWallpaper", 0, win32con.REG_SZ, "0")  # 修改是否平铺：值为0或1,标识是否要像瓦片一样把桌面铺起来
    win32gui.SystemParametersInfo(win32con.SPI_SETDESKWALLPAPER, imagePath, 1 + 2)  # 设置桌面壁纸，指定路径图片
    print("已设置桌面壁纸，图片为：" + imagePath)


# 每天获取不同的桌面壁纸
def get_imag_path(dirPath):
    global emptyPic
    if not os.path.exists(dirPath):
        print("文件夹不存在")
        return ''

    today = datetime.today().isoweekday()
    today = str(today)
    if today == '1':
        change_imag_sort(dirPath)

    for root, dirs, files in os.walk(dirPath):
        for file in files:
            emptyPic = file
            if file.find(str(today)) != -1:
                return os.path.join(root, file)
    return emptyPic


# 每周一打乱图片顺序，重新设置图片名称
def change_imag_sort(dirPath):
    # 先重置名称，防止名称重复
    new_name = 'wait_change'
    new_num = 0
    for file in os.listdir(dirPath):
        suffix = "." + file.split(".")[1]
        old_filename = dirPath + os.sep + file
        new_file_name = dirPath + os.sep + new_name + str(new_num) + suffix
        os.rename(old_filename, new_file_name)
        new_num += 1

    # 再设置确切的名称
    name = '高达'
    new_list = [i for i in range(1, 8)]
    for file in os.listdir(dirPath):
        suffix = "." + file.split(".")[1]
        list_len = len(new_list)
        random_num = random.randint(0, list_len - 1)
        num = new_list[random_num]
        old_filename = dirPath + os.sep + file
        new_file_name = dirPath + os.sep + name + str(num) + suffix
        os.rename(old_filename, new_file_name)
        new_list.remove(num)
    print("成功随机修改文件名称")


if __name__ == '__main__':
    dirPath = 'D:\测试专用\wallpaper'
    imagePath = get_imag_path(dirPath)
    if imagePath:
        set_wallpaper(imagePath)
    else:
        print("找不到对应的图片，无法设置桌面壁纸")

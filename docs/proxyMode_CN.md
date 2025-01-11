### 模式2

> 🌐️ 中文 | [English](proxyMode.md)

> 注意，信任证书是一件非常危险的事情，如果您将代理设置为系统代理，代理服务能够能够捕获您电脑的所有网络请求，因此对于不信任的程序，千万不要信任证书并将代理设置为系统代理。

> 模式2是 cursor-vip 在您电脑本地启动一个代理服务，只将 cursor 客户端的代理配置指向 cursor-vip 的代理服务，其他软件不走此代理，我们承诺不会记录您的任何网络请求。

如果选择模式2首次启动后需安装信任证书，证书会在首次启动命令后自动生成，路径为 `~/.cursor-vip/`。

#### MacOS: 
命令行方式
```sh
sudo security add-trusted-cert -d -p ssl -p basic -k /Library/Keychains/System.keychain ~/.cursor-vip/i-need-to-trust-ca-cert.pem
```
图形化方式
> 在终端执行 `open ~/.cursor-vip`，双击 i-need-to-trust-ca-cert.pem 文件，选择`登录`，弹出“钥匙串访问”窗口，选择证书，搜索 cursor-vip，双击 cursor-vip，展开信任，选择 X.509基本策略 “始终信任”，关闭弹窗，输入密码确认，证书安装完成。
#### Windows: 
命令行方式（gitBash）
```sh
certutil -addstore root ~/.cursor-vip/ca-cert.cer
```
图形化方式
> 在windows搜索输入 `管理用户证书`,选择`管理用户证书`功能，展开`受信任的根证书颁发机构`，选中`证书`，右键`所有任务`，选择`导入`，下一步，输入`%homepath%\.cursor-vip\ca-cert.cer`文件 (如果用选择方式的话，.pem会被隐藏，右下角选择*可以看到)，一直下一步，完成;。
#### Linux
> [Ubuntu/Debian](https://askubuntu.com/questions/73287/how-do-i-install-a-root-certificate/94861#94861)、[Fedora](https://docs.fedoraproject.org/en-US/quick-docs/using-shared-system-certificates/#proc_adding-new-certificates)

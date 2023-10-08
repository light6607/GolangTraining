#!/bin/bash
# use to fix vscode remote ssh download server error
echo "ps:only support x64 linux case,other arch system change the scripts"
read -p "please input vscode的commit id : " commit
if [ ! -z "$commit" ]; then
  test -d /root/.vscode-server/bin/$commit || mkdir -p /root/.vscode-server/bin/$commit
  echo "Downloading vscode package..."
  mkdir -p ~/vscode-backup
  wget https://vscode.cdn.azure.cn/stable/$commit/vscode-server-linux-x64.tar.gz -O ~/vscode-backup/vscode-server-linux-x64.tar.gz
  tar -zxvf ~/vscode-backup/vscode-server-linux-x64.tar.gz -C  ~/vscode-backup
  cp -rf ~/vscode-backup/vscode-server-linux-x64/* /root/.vscode-server/bin/$commit/  
  echo "Vscode server configure success"
  rm -rf ~/vscode-backup/
  echo "Clean up temporary files"
else
    echo "Please input vscode version commit id,see it in Abount Visual Code"
fi
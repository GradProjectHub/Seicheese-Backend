#!/bin/bash

# 証明書の保存ディレクトリ作成
mkdir -p ssl

# 一時的にGoサーバーを停止（ポート解放）
echo "Stopping Go server..."
docker compose stop go

# 証明書の取得
echo "Getting SSL certificate..."
docker compose run --rm certbot

# 証明書が正常に取得できたか確認
if [ $? -eq 0 ]; then
    echo "Certificate obtained successfully"
    
    # 証明書のコピー
    echo "Copying certificates..."
    docker compose run --rm certbot cp -L /etc/letsencrypt/live/api.seicheese.jp/fullchain.pem /ssl/
    docker compose run --rm certbot cp -L /etc/letsencrypt/live/api.seicheese.jp/privkey.pem /ssl/

    # 権限の設定
    echo "Setting permissions..."
    sudo chown -R $USER:$USER ssl/
    chmod 600 ssl/*.pem

    # Goサーバーの再起動
    echo "Restarting Go server..."
    docker compose start go
    
    echo "SSL setup completed successfully"
else
    echo "Failed to obtain certificate"
    docker compose start go
    exit 1
fi
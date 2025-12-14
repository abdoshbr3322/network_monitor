APP_NAME="netmon"
DAEMON_NAME="netmond"
APP_USER="netmon"

go mod tidy

go build -o $APP_NAME ./cmd/app
go build -o $DAEMON_NAME ./cmd/daemon

sudo mv $APP_NAME /usr/bin
sudo mv $DAEMON_NAME /usr/bin

sudo useradd --system --no-create-home --shell /usr/sbin/nologin $APP_USER

sudo mkdir /var/lib/$APP_NAME
sudo chown $APP_NAME:$APP_NAME /var/lib/$APP_NAME

sudo cp ./systemd/$APP_NAME.service /etc/systemd/system

sudo systemctl daemon-reexec
sudo systemctl daemon-reload

sudo systemctl enable "$APP_NAME.service"
sudo systemctl start "$APP_NAME.service"

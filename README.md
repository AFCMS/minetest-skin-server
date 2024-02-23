# Minetest Skin Server

![GitHub Workflow Status](https://img.shields.io/github/checks-status/AFCMS/minetest-skin-server/master?style=flat-square)

> [!IMPORTANT]
> This server is still in development and is not ready for production use.
> Breaking changes may occur at any time.

This server is made for serving Minetest skins to Minetest servers. It is licensed under GPLv3.

-   ✅ Easy to use and powerful **API**
-   ✅ Skins compatible with both **MineClone2** and **Minetest Game**
-   ✅ Fast and reliable, thanks to **Docker**, **Golang**, **Fiber** and **PostgreSQL**
-   ✅ Optimised images using **OptiPNG**

## Design

The server is build with the **Go** language on-top of the [**Fiber**](https://gofiber.io/) framework.

It uses also the [**GORM**](https://gorm.io) library for interacting with the database.

The frontend is build with the [**React**](https://reactjs.org) library and the following modules:

-   [**TailwindCSS**](https://tailwindcss.com) for styling
-   [**HeadlessUI**](https://headlessui.com) for dialogs, combobox, etc
-   [**Heroicons**](https://heroicons.com) for most icons
-   [**React Router**](https://reactrouter.com)
-   [**React Three Fiber**](https://github.com/pmndrs/react-three-fiber) for the 3D preview of skins
-   [**Recoil**](https://recoiljs.org) for state managment

## Running Server

### Production

#### 1. Install `docker`

Follow the official guide for your OS.

- [Ubuntu](https://docs.docker.com/engine/install/ubuntu)
- [Debian](https://docs.docker.com/engine/install/debian)
- [Fedora](https://docs.docker.com/engine/install/fedora)
- [RHEL/CentOS](https://docs.docker.com/engine/install/centos)

> [!NOTE]
> The installation links are from Docker Engine, which works only under Linux.
>
> [Docker Desktop](https://www.docker.com/products/docker-desktop) can be used on Windows, MacOS and Linux.
>
> It runs a Linux VM in the background and isn't as performant as the native version, but it's easier to install and use.

#### 2. Download source code

```shell
git clone https://github.com/AFCMS/minetest-skin-server && cd minetest-skin-server
```

#### 3. Configure server

```shell
cp exemple.env .env
```

Edit the `.env` file with the config you want.

A typical production config would be:

```ini
JWT_SECRET=secret
DEBUG_DATABASE=false
ENABLE_OPTIPNG=true

DB_HOST=db
DB_USER=user
DB_PASSWORD=azerty
DB_PORT=5432
DB_NAME=skin_server
```

#### 4. Run service

```shell
docker compose up --build
```

> [!NOTE]
> In case you do not want to build the image yourself, you can use the one present on the GitHub registery.
>
> ```py
> docker pull ghcr.io/afcms/minetest-skin-server:master
> ```
>
> You will be able to use it just by removing the `--build` flag from the docker-compose command line.

### Development

It may be easier to not use docker while developing, both for frontend and backend development.

#### 1. Install Go and NodeJS

Follow the official guides for you OS.

I recommend using NodeJS v20 installed using [**nvm**](https://github.com/nvm-sh/nvm) under linux.

#### 2. (Optional) Install OptiPNG

If you want to enable [OptiPNG](https://optipng.sourceforge.net), you need to install it.

On Ubuntu/Debian:

```shell
sudo apt install optipng
```

On Fedora/RHEL:

```shell
sudo dnf install optipng
```

#### 3. Download source code

```shell
git clone https://github.com/AFCMS/minetest-skin-server && cd minetest-skin-server
```

#### 4. Install Go dependencies

```shell
go mod download
```

#### 5. Install NodeJS dependencies

```shell
cd frontend && npm install --include=dev && cd ..
```

#### 6. Configure server

```shell
cp exemple.env .env
```

Edit the `.env` file with the config you want.

A typical development config would be:

```ini
JWT_SECRET=secret
DEBUG_DATABASE=true
ENABLE_OPTIPNG=true

DB_HOST=db
DB_USER=user
DB_PASSWORD=azerty
DB_PORT=5432
DB_NAME=skin_server
```

#### 7. Build frontend

The frontend served by the Fiber backend needs to be build before running the app.

This can be done like this:

```shell
cd frontend && npm run build && cd ..
```

> [!NOTE]
> You can run the frontend in development mode on another port with `npm start` while the backend is running, it will still use the backend's API

#### 8. Build and run backend

```shell
go build && ./minetest-skin-server
```

## Development Environment

I recommand using either **VSCode** or **GoLand** with the **Prettier** autoformater.

There are multiple VSCode extensions marked as recommended for the workspace.

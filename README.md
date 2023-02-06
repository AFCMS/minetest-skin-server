# Minetest Skin Server

![GitHub Workflow Status](https://img.shields.io/github/checks-status/AFCMS/minetest-skin-server/master?style=flat-square)

> ⚠️ This server is still in development and is not ready for production use.
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

1. Install `docker` and `docker-compose`

Under Debian/Ubuntu:

```sh
sudo apt install docker docker-compose
```

2. Download source code

```sh
git clone https://github.com/AFCMS/minetest-skin-server && cd minetest-skin-server
```

3. Configure server

```sh
cp exemple.env .env
```

Edit the `.env` file with the config you want.

A typical production config would be:

```ini
USE_SQLITE=false
DEBUG_DATABASE=false
JWT_SECRET=secret
ENABLE_OPTIPNG=true
POSTGRES_DB=skin_server
POSTGRES_PASSWORD=adminpassword
POSTGRES_USER=admin
```

4. Run service

```sh
docker-compose up --build
```

> **Note** In case you do not want to build the image yourself, you can use the one present on the GitHub registery.
>
> ```py
> docker pull ghcr.io/afcms/minetest-skin-server:master
> ```
>
> You will be able to use it just by removing the `--build` flag from the docker-compose command line.

### Development

It may be easier to not use docker while developing, both for frontend and backend development.

1. Install Go and NodeJS

Follow the official guides for you OS.

I recommend using NodeJS v16 installed using [**nvm**](https://github.com/nvm-sh/nvm) under linux.

2. (Optional) Install OptiPNG

If you want to enable OptiPNG, you need to install it.

On Ubuntu, it is as easy as:

```sh
sudo apt install optipng
```

3. Download source code

```sh
git clone https://github.com/AFCMS/minetest-skin-server && cd minetest-skin-server
```

4. Install Go dependencies

```sh
go mod download
```

5. Install NodeJS dependencies

```sh
cd frontend && npm install --include=dev && cd ..
```

6. Configure server

```sh
cp exemple.env .env
```

Edit the `.env` file with the config you want.

> **Warning**
> If you don't want to setup a PostgreSQL database, you can use SQLite instead
>
> It will use a `database.db` file created at the root directory of the app

A typical development config would be:

```ini
USE_SQLITE=true
DEBUG_DATABASE=true
JWT_SECRET=secret
ENABLE_OPTIPNG=true
POSTGRES_DB=
POSTGRES_PASSWORD=
POSTGRES_USER=
```

7. Build frontend

The frontend served by the Fiber backend needs to be build before running the app.

This can be done like this:

```sh
cd frontend && npm run build && cd ..
```

> **Note**
> You can run the frontend in development mode on another port with `npm start` while the backend is running, it will still use the backend's API

8. Build and run backend

```sh
go build && ./minetest-skin-server
```

## Development Environment

I recommand using either **VSCode** or **GoLand** with the **Prettier** autoformater.

There are multiple VSCode extensions marked as recommended for the workspace.

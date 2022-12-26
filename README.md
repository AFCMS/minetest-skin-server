# Minetest Skin Server

This server is made for serving Minetest skins to Minetest servers.

## Running Server

> A Docker configuration is planed, but still not functional

Use the following commands to run the backend:

```sh
git clone https://github.com/AFCMS/minetest-skin-server.git

cd minetest-skin-server

go mod download

go build && ./minetest-skin-server
```

The frontend isn't linked to the backend yet so you have to run it separatly:

```sh
cd frontend

npm install --dev

npm start
```

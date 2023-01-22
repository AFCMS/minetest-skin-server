-----------------------------------
--- Minetest Skin Server Client ---
--- by AFCMS, under GPLv3       ---
-----------------------------------

---@class MSSSkinDetails
---@field description string The description of the skin
---@field public boolean
---@field owner_id string ID of the owner
---@field created_at string The creation date

---@class MSSClient
---@field url string URL of the server
---@field http_client HTTPApiTable
local client_api = {}

---Fetch skin details
---@param uuid string
---@param callback fun(sucess: boolean, details?: MSSSkinDetails)
function client_api:fetch_skin(uuid, callback)
	print(self.url)
	self.http_client.fetch({
		method = "GET",
		url = self.url .. "/api/skin/skin/" .. uuid,
	}, function(res)
		if res.succeeded == true then
			local data = minetest.parse_json(res.data)
			if type(data) ~= "table" then
				callback(false, nil)
				return
			end
			callback(true, {
				description = data.description,
				public = data.public,
				owner_id = data.owner_id,
				created_at = data.created_at
			})
		else
			callback(false, nil)
		end
	end)
end

---Create a new API client
---@param url string URL of the server
---@param http_client HTTPApiTable The HTTP Client
---@return MSSClient
return function(url, http_client)
	return setmetatable({url = url, http_client = http_client}, {__index = client_api})
end

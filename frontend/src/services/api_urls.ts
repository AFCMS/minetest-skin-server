import isDev from "./is_dev";

const prefix = isDev() ? "http://localhost:8080" : "";

// Contains the API URLs
const ApiUrls = {
	Info: prefix + "/api/info",
	AccountLogin: prefix + "/api/account/login",
	AccountLogout: prefix + "/api/account/logout",
	AccountRegister: prefix + "/api/account/register",
	AccountUser: prefix + "/api/account/user",
	SkinList: prefix + "/api/skin/list",
	SkinCreate: prefix + "/api/skin/create",
	SkinRecent: prefix + "/api/skin/recent",
};

export default ApiUrls;

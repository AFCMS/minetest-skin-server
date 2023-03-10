import { atom, selector } from "recoil";

const AuthState = atom({
	key: "AuthState",
	default: {
		id: undefined,
		email: undefined,
		name: undefined,
	},
});

const AuthStateEmail = selector({
	key: "AuthStateEmail",
	get: ({ get }) => {
		return get(AuthState).email;
	},
	set: ({ set, get }) => {},
});

const AuthStateIsAuthenticated = selector({
	key: "AuthStateIsAuthenticated",
	get: ({ get }) => {
		let a = get(AuthState);
		return (
			a.email !== undefined || a.id !== undefined || a.name !== undefined
		);
	},
});

export { AuthState, AuthStateEmail, AuthStateIsAuthenticated };

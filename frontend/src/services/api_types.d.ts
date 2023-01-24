declare namespace ApiTypes {
	type UserResponse = {
		id: number;
		name: string;
		email: string;
		permission_level: 1 | 2 | 3 | 4;
		banned?: boolean;
		ban_reason?: string;
		created_at: string;
		last_connection: string;
	};

	type SkinResponse = {
		uuid: string;
		description: string;
		public?: boolean;
		approved?: boolean;
		owner_id: number;
		creation_date: string;
	};
}

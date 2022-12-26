function Register() {
	return (
		<div className="justify-center align-middle flex">
			<div className="flex flex-col max-w-md p-6 gap-4 m-10 app-pannel">
				<h1 className="panel-text-heading">Register</h1>
				<label className="block">
					<span className="text-slate-800">Username</span>
					<input
						type="text"
						name="username"
						id="username"
						placeholder="GamerPro"
						className="form"
					/>
				</label>
				<label className="block">
					<span className="text-slate-800">Email</span>
					<input
						type="email"
						name="email"
						id="email"
						placeholder="Ex: someone@gmail.com"
						className="form"
					/>
				</label>
				<label className="block">
					<span className="text-slate-800">Password</span>
					<input
						type="password"
						name="password"
						id="password"
						placeholder="SuperStrongPassword"
						className="form"
					/>
				</label>

				<div className="align-baseline mt-2">
					<input
						type="checkbox"
						name="understand"
						id="understand"
						className="form-checkbox"
						defaultChecked={false}
					/>
					<span className="ml-2 text-slate-800 text-sm">
						I aggree to the TOS
					</span>
				</div>
				<button className="button-primary w-full mt-2">Register</button>
			</div>
		</div>
	);
}

export default Register;

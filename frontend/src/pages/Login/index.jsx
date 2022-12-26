function Login() {
	return (
		<div className="justify-center align-middle flex">
			<div className="flex flex-col max-w-md p-6 gap-4 m-10 app-pannel">
				<h1 className="panel-text-heading">Login</h1>
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
				<div className="flex gap-4 mt-4">
					<button className="button-secondary w-full">
						Register
					</button>
					<button className="button-primary w-full">Login</button>
				</div>
			</div>
		</div>
	);
}

export default Login;

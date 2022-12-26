function About() {
	return (
		<div className="justify-center align-middle flex">
			<div className=" bg-slate-50 m-10 rounded shadow-md container p-8 max-w-prose">
				<h1 className="panel-text-heading">About</h1>
				<p className="mt-2">
					Minetest Skin Server is a service to upload Minetest skins
					and let servers retreive them. It is licensed under GPLv3.
				</p>
			</div>
		</div>
	);
}

export default About;

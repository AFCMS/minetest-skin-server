import { useRef } from "react";
import { NearestFilter } from "three";
import { useGLTF, useTexture } from "@react-three/drei";
import skinModel from "../../assets/skin_character.gltf";
import skinTexture from "../../assets/character_test.png";

function SkinModel() {
	const gltf = useGLTF(skinModel);
	const texture = useTexture(skinTexture);

	// Make texture render pixelated
	texture.magFilter = NearestFilter;

	const myMesh = useRef();
	return (
		<mesh
			ref={myMesh}
			receiveShadow
			castShadow
			geometry={gltf.nodes.Player.geometry}
			material={gltf.materials.Character}
		>
			<meshStandardMaterial map={texture} />
		</mesh>
	);
}

// Model can be preloaded, since its the same for all skins
useGLTF.preload(skinModel);

export default SkinModel;

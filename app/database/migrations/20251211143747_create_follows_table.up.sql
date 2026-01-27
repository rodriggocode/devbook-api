CREATE TABLE follows(
	user_id INT NOT NULL,
        follower_id INT NOT NULL,

	-- modelo composto, entao essas duas vao ser a chave primaria
	PRIMARY KEY (user_id, follower_id),
	
	FOREIGN KEY (user_id) REFERENCES users(id_user) ON DELETE CASCADE,
        FOREIGN KEY (follower_id) REFERENCES users(id_user) ON DELETE CASCADE,
 	
	UNIQUE (user_id, follower_id)
);

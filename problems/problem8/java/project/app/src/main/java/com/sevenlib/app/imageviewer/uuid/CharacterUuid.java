package com.sevenlib.app.imageviewer.uuid;

public class CharacterUuid extends Uuid {
	public CharacterUuid() {
		super.type = UuidType.CHARACTER;
		super.data = "null";
	}
	
	public CharacterUuid(String uuid) {
		super.type = UuidType.CHARACTER;
		super.data = uuid;
	}
	
	public void set(Object obj) {
		super.data = obj;
	}

	public Object get() {
		return super.data;
	}
	
	@Override
	public boolean equals(Object obj) {
		return false;  
	}
}

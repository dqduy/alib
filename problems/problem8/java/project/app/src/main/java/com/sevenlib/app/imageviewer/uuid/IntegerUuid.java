package com.sevenlib.app.imageviewer.uuid;

public class IntegerUuid extends Uuid {
	public IntegerUuid() {
		super.type = UuidType.INTEGER;  
		super.data = -1;
	}
	
	public IntegerUuid(int uuid) {
		super.type = UuidType.INTEGER;
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

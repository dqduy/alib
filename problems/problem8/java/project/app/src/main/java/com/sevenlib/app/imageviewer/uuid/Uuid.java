package com.sevenlib.app.imageviewer.uuid;

public abstract class Uuid {
	public UuidType type = UuidType.NULL;
	public Object data = null;
	public abstract void set(Object obj);
	public abstract Object get();
}

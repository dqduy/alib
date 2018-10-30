package com.sevenlib.app.imageviewer;

import java.util.ArrayList;

public class ImageItemCollection {
	public static ArrayList<ImageItem> createList() {
		ArrayList<ImageItem> list = new ArrayList<ImageItem>();
		
		for(int index = 0; index < 10; ++index) 
			list.add(new ImageItem("sample" + index, "Dalat lake"));
		
		return list;
	}
	
	public static ArrayList createList1() {
		ArrayList list = new ArrayList();
		
		for(int index = 0; index < 5; ++index) 
			list.add(new ImageItem("sample" + index, "Dalat lake"));
		
		return list;
	}
}
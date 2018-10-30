package com.sevenlib.app.imageviewer;

import android.util.Log;

public class PLog {
	public static final String MAIN_TAG = "duy";

	public static void WriteLog(String tag, String content) {
		Log.i(tag, content);
	}
}
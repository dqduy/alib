package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.support.v7.widget.RecyclerView;
import android.support.v7.widget.LinearLayoutManager;
import android.widget.TextView;

import java.util.ArrayList;

public class MainAct extends Activity {
    static {
        System.loadLibrary("hello");
    }

    private native String getMessage();

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        android.util.Log.i("duy", "Begin app");
        setContentView(R.layout.list);
		
		ArrayList<ImageItem> list = ImageItemCollection.createList();
        RecyclerView listView = (RecyclerView) findViewById(R.id.my_list);
        ImageItemsAdapter adapter = new ImageItemsAdapter(list);
        listView.setLayoutManager(new LinearLayoutManager(this));
        listView.setAdapter(adapter);
    }
}

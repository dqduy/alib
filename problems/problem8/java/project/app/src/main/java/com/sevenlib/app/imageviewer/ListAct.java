package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.ListActivity;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.TextView;
import android.view.View;
import 	android.content.res.Resources;

import java.util.ArrayList;

public class ListAct extends ListActivity {
    private TextView selection;
    
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.list_main);
        Resources res = getResources();

        String[] planets = res.getStringArray(R.array.planets);
        ArrayAdapter<String> adapter = new ArrayAdapter<String>(this, 
                                        android.R.layout.simple_list_item_1,
                                        planets);
        setListAdapter(adapter);

        selection = (TextView) findViewById(R.id.selection);
    }

    @Override
    public void onListItemClick(ListView parent, View v, int pos, long id) {
        selection.setText(pos);
    }
}

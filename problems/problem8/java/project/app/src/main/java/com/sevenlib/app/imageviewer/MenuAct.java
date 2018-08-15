<<<<<<< HEAD
/*
    Author: Duy Quoc
    Title: Introduction Android life cycle
*/

=======
>>>>>>> 1a2623ba1a0a2575a1a9fb484594cf40bd1182a6
package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
<<<<<<< HEAD
import android.content.SharedPreferences;
import android.widget.TextView;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Toast;

public class MenuAct extends AppCompatActivity {

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.menu_act);

        Toolbar bar = (Toolbar) findViewById(R.id.my_toolbar);
        setSupportActionBar(bar);
        //getSupportActionBar().setDisplayShowHomeEnabled(true);
        //getSupportActionBar().setHomeAsUpIndicator(R.drawable.ic_menu);
        bar.setNavigationIcon(R.drawable.horizontal_lines_32);
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.search) {
            Toast.makeText(MenuAct.this, "Search clicked", Toast.LENGTH_LONG).show();
            return true;
        }
        else if (id == R.id.settings) {
            Toast.makeText(MenuAct.this, "Settings clicked", Toast.LENGTH_LONG).show();
            return true;
        }

        return super.onOptionsItemSelected(item);
=======
import android.support.v7.widget.RecyclerView;
import android.support.v7.widget.LinearLayoutManager;
import android.widget.TextView;

import java.util.ArrayList;

public class MenuAct extends Activity {
    static {
        System.loadLibrary("hello");
    }

    private native String getMessage();

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

>>>>>>> 1a2623ba1a0a2575a1a9fb484594cf40bd1182a6
    }
}

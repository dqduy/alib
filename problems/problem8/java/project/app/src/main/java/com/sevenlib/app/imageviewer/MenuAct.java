/*
    Author: Duy Quoc
    Title: Introduction Android life cycle
*/

package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.widget.TextView;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.support.v7.app.ActionBar;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Toast;
import android.support.v4.widget.DrawerLayout;
import android.support.v4.view.GravityCompat;
import android.support.design.widget.NavigationView;
import android.support.v7.app.ActionBarDrawerToggle;

public class MenuAct extends AppCompatActivity {
    private DrawerLayout drawer;
    private ActionBarDrawerToggle t;

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.menu_act);

        Toolbar bar = (Toolbar) findViewById(R.id.my_toolbar);
        setSupportActionBar(bar);
        ActionBar actionbar = getSupportActionBar();
        actionbar.setDisplayHomeAsUpEnabled(true);
        actionbar.setHomeAsUpIndicator(R.drawable.horizontal_lines_24);

        drawer = (DrawerLayout) findViewById(R.id.drawer_layout);

        NavigationView nav = (NavigationView) findViewById(R.id.nav_view);
        nav.setNavigationItemSelectedListener(
            new NavigationView.OnNavigationItemSelectedListener() {
                @Override
                public boolean onNavigationItemSelected(MenuItem menuItem) {
                    // set item as selected to persist highlight
                    menuItem.setChecked(true);
                    // close drawer when item is tapped
                    drawer.closeDrawers();

                    // Add code here to update the UI based on the item selected
                    // For example, swap UI fragments here

                    return true;
                }
            });
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

        PLog.WriteLog(PLog.MAIN_TAG, "item: " + id);
    
        if (id == R.id.search) {
            Toast.makeText(MenuAct.this, "Search clicked", Toast.LENGTH_LONG).show();
            return true;
        }
        else if (id == R.id.settings) {
            Toast.makeText(MenuAct.this, "Settings clicked", Toast.LENGTH_LONG).show();
            return true;
        }
        else if (id == android.R.id.home) {
            drawer.openDrawer(GravityCompat.START);
            return true;
        }

        return super.onOptionsItemSelected(item);

    }
}
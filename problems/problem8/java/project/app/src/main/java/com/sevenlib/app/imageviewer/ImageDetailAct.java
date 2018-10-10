/*
    Author: Duy Quoc
    Title: Display a image in one activity
*/

package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.widget.ImageView;
import android.support.v7.app.AppCompatActivity;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.support.v7.app.ActionBar;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Toast;
import android.support.v7.widget.LinearLayoutManager;

import java.lang.CharSequence;
import java.util.ArrayList;

public class ImageDetailAct extends AppCompatActivity {

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.image_detail);

        CharSequence title = "Doraemon 1, page 21";
        Toolbar bar = (Toolbar) findViewById(R.id.my_toolbar);
        setSupportActionBar(bar);
        ActionBar actionbar = getSupportActionBar();
        actionbar.setDisplayHomeAsUpEnabled(true);
        actionbar.setHomeAsUpIndicator(R.drawable.left_arrow);
        actionbar.setTitle(title);
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.image_settings, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        PLog.WriteLog(PLog.MAIN_TAG, "item: " + id);
    
        switch(id) {
            case android.R.id.home:
                Toast.makeText(ImageDetailAct.this, "Home clicked", Toast.LENGTH_SHORT).show();
                break;
            case R.id.share:
                Toast.makeText(ImageDetailAct.this, "Share clicked", Toast.LENGTH_SHORT).show();
                break;
            case R.id.delete:
                Toast.makeText(ImageDetailAct.this, "Delete clicked", Toast.LENGTH_SHORT).show();
                break;
            case R.id.settings:
                //Toast.makeText(ImageDetailAct.this, "Settings clicked", Toast.LENGTH_SHORT).show();
                break;                
            default:
                Toast.makeText(ImageDetailAct.this, "No one clicked", Toast.LENGTH_SHORT).show();
                break;
        }

        return super.onOptionsItemSelected(item);
    }
}
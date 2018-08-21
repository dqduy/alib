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
        actionbar.setHomeAsUpIndicator(R.drawable.left_arrow_128);
        actionbar.setTitle(title);
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.image_settings, menu);
        return true;
    }

}
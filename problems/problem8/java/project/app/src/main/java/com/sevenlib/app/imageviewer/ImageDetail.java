/*
    Author: Duy Quoc
    Title: Display a image in one activity
*/

package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.widget.ImageView;
import android.support.v7.app.AppCompatActivity;
import java.util.ArrayList;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import java.io.InputStream;

public class ImageDetail extends AppCompatActivity {

    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.image_detail);

        ImageView view = (ImageView) findViewById(R.id.pic1);

        BitmapFactory.Options opt = new BitmapFactory.Options();
        //InputStream input = getResources().openRawResource(R.raw.london);
        //opt.inJustDecodeBounds = true;
        opt.inScaled = false;
        opt.inSampleSize = 2;
        opt.inPreferredConfig = Bitmap.Config.ALPHA_8;
        Bitmap bmp = BitmapFactory.decodeResource(getResources(), R.raw.london, opt);
        //int outWidth = opt.outWidth;
        //int outHeight = opt.outHeight;
        //String type = opt.outMimeType;
        //PLog.WriteLog(PLog.MAIN_TAG, "width: " + outWidth + ", height: " + outHeight + ", type: " + type);
    
        view.setImageBitmap(bmp);
        PLog.WriteLog(PLog.MAIN_TAG, "width: " + bmp.getWidth() + ", height: " + bmp.getHeight());
        PLog.WriteLog(PLog.MAIN_TAG, "Byte count: " + bmp.getByteCount());
        PLog.WriteLog(PLog.MAIN_TAG, "Allo count: " + bmp.getAllocationByteCount());
        PLog.WriteLog(PLog.MAIN_TAG, "Config: " + bmp.getConfig().toString());
    }



}
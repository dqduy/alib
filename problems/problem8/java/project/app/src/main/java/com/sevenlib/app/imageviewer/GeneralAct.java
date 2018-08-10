/*
    Author: Duy Quoc
    Title: Introduction Android life cycle
*/

package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.util.Log;
import android.view.Window;  
import java.io.File;
import android.os.Environment;

public class GeneralAct extends Activity {
    private final String TAG_NAME = "duy";
    //Pair 1
    @Override
    public void onCreate(Bundle savedInstanceState) {
        PLog.WriteLog(PLog.MAIN_TAG, "onCReate()");

        this.requestWindowFeature(Window.FEATURE_NO_TITLE);
        super.onCreate(savedInstanceState);

        getPathFromDevice();
    }

    private void getPathFromDevice() {

        File file = this.getFilesDir();
        PLog.WriteLog(PLog.MAIN_TAG, "=========BEGIN Internal Storage====================" );
        PLog.WriteLog(PLog.MAIN_TAG, "getName(): " + file.getName());
        PLog.WriteLog(PLog.MAIN_TAG, "getPath():         " + file.getPath());
        PLog.WriteLog(PLog.MAIN_TAG, "getAbsolutePath(): " + file.getAbsolutePath());
        PLog.WriteLog(PLog.MAIN_TAG, "=========END Internal Storage====================" );

        file = this.getExternalFilesDir(null);
        PLog.WriteLog(PLog.MAIN_TAG, "=========BEGIN Private - External Storage====================" );
        PLog.WriteLog(PLog.MAIN_TAG, "getName(): " + file.getName());
        PLog.WriteLog(PLog.MAIN_TAG, "getPath():         " + file.getPath());
        PLog.WriteLog(PLog.MAIN_TAG, "getAbsolutePath(): " + file.getAbsolutePath());
        PLog.WriteLog(PLog.MAIN_TAG, "=========END Private - External Storage=====================" );

        file = Environment.getExternalStoragePublicDirectory (Environment.DIRECTORY_PICTURES);
        PLog.WriteLog(PLog.MAIN_TAG, "=========BEGIN Public - External Storage====================" );
        PLog.WriteLog(PLog.MAIN_TAG, "getName(): " + file.getName());
        PLog.WriteLog(PLog.MAIN_TAG, "getPath():         " + file.getPath());
        PLog.WriteLog(PLog.MAIN_TAG, "getAbsolutePath(): " + file.getAbsolutePath());
        PLog.WriteLog(PLog.MAIN_TAG, "=========END Public - External Storage=====================" );

        file = Environment.getExternalStorageDirectory ();
        PLog.WriteLog(PLog.MAIN_TAG, "=========BEGIN Primary - External Storage====================" );
        PLog.WriteLog(PLog.MAIN_TAG, "getName(): " + file.getName());
        PLog.WriteLog(PLog.MAIN_TAG, "getPath():         " + file.getPath());
        PLog.WriteLog(PLog.MAIN_TAG, "getAbsolutePath(): " + file.getAbsolutePath());
        PLog.WriteLog(PLog.MAIN_TAG, "=========END Primary - External Storage=====================" );

        File[] list = this.getExternalMediaDirs();
        int length  = list.length;
        PLog.WriteLog(PLog.MAIN_TAG, "List sdcard: " + length);

        for(int index = 0; index < length; ++index) {
            file = list[index];
            PLog.WriteLog(PLog.MAIN_TAG, "=========BEGIN External Storage==================== " + index);
            PLog.WriteLog(PLog.MAIN_TAG, "getName(): " + file.getName());
            PLog.WriteLog(PLog.MAIN_TAG, "getPath():         " + file.getPath());
            PLog.WriteLog(PLog.MAIN_TAG, "getAbsolutePath(): " + file.getAbsolutePath());
            PLog.WriteLog(PLog.MAIN_TAG, "=========END   External Storage=====================" );
        }


    }

    @Override
    public void onDestroy() {
        Log.i("duy", "onDestroy()");
        super.onDestroy();
    }

    //Pair 2    
    @Override
    public void onStart() {
        Log.i("duy", "onStart()");
        super.onStart();
    }

    @Override
    public void onStop() {
        Log.i("duy", "onStop()");
        super.onStop();
    }

    //Pair 3
    @Override
    public void onResume() {
        Log.i("duy", "onResume()");
        super.onResume();
    }

    @Override
    public void onPause() {
        Log.i("duy", "onPause()");
        super.onPause();
    }

    //Remain
    @Override
    public void onRestart() {
        Log.i("duy", "onRestart()");
        super.onRestart();
    }
}

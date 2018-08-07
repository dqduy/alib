/*
    Author: Duy Quoc
    Title: Introduction Android life cycle
*/

package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
import android.util.Log;

public class GeneralAct extends Activity {
    
    //Pair 1
    @Override
    public void onCreate(Bundle savedInstanceState) {
        Log.i("duy", "onCreate()");
        super.onCreate(savedInstanceState);        
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

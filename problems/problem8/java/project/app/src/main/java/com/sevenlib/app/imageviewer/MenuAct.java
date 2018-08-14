package com.sevenlib.app.imageviewer;

import android.os.Bundle;
import android.app.Activity;
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

    }
}

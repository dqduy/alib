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
import android.content.Intent;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.FileOutputStream;
import java.io.IOException;
import android.net.Uri;
import android.support.v4.provider.DocumentFile;

public class GeneralAct extends Activity {
    private final String TAG_NAME = "duy";
    //Pair 1
    @Override
    public void onCreate(Bundle savedInstanceState) {
        PLog.WriteLog(PLog.MAIN_TAG, "onCReate()");

        this.requestWindowFeature(Window.FEATURE_NO_TITLE);
        super.onCreate(savedInstanceState);

        Intent sender = new Intent(Intent.ACTION_OPEN_DOCUMENT_TREE);
        startActivityForResult(sender, 25);
        getPathFromDevice();

        File file = new File("/storage/external_SD","lg.txt");
        StringBuilder text = new StringBuilder();
        try {
            BufferedReader br = new BufferedReader(new FileReader(file));
            String line;
        
            while ((line = br.readLine()) != null) {
                text.append(line);
                text.append('\n');
            }
            br.close();

            //PLog.WriteLog(PLog.MAIN_TAG, text.toString());
        }
        catch (Exception e) {
            //You'll need to add proper error handling here
        }

        {
            File out = new File("/storage/external_SD","lg1.txt");
            
            String data = "result removable";
            try {
                FileOutputStream outStream = new FileOutputStream(out);
                outStream.write(data.getBytes());
                outStream.close();
            }
            catch (Exception e) {
                PLog.WriteLog(PLog.MAIN_TAG, "Fail to write to removable storage");
            } 
            finally {
                
            }
        }

        {
            File out = new File("/storage/emulated/0","lg1.txt");
            
            String data = "result external";
            try {
                FileOutputStream outStream = new FileOutputStream(out);
                outStream.write(data.getBytes());
                outStream.close();
            }
            catch (Exception e) {
                PLog.WriteLog(PLog.MAIN_TAG, "Fail to write to external storage");
            } 
            finally {
                
            }
        }
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (requestCode == 25 && resultCode == RESULT_OK) {
            //getContentResolver().takePersistableUriPermission(data.getData(), Intent.FLAG_GRANT_READ_URI_PERMISSION | Intent.FLAG_GRANT_WRITE_URI_PERMISSION);
            Uri treeUri = data.getData();
            PLog.WriteLog(PLog.MAIN_TAG, "Uri: " + treeUri.toString());    
            PLog.WriteLog(PLog.MAIN_TAG, "Uri - getPath(): " + treeUri.getPath());
            DocumentFile docRoot = DocumentFile.fromTreeUri(this, treeUri);
            PLog.WriteLog(PLog.MAIN_TAG, "Found item: " + docRoot.getName() + " - " + docRoot.getType());
            for(DocumentFile file: docRoot.listFiles()) {
                if(file.isFile())
                    PLog.WriteLog(PLog.MAIN_TAG, "Found item: file" + file.getName() + " - " + file.getType());
                else
                    PLog.WriteLog(PLog.MAIN_TAG, "Found item: dir " + file.getName() + " - " + file.getType());
            }

        } 
        super.onActivityResult(requestCode, resultCode, data);
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

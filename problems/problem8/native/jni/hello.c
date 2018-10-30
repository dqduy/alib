#include <jni.h>
#include "include/hello.h"
#include <android/log.h> 

JNIEXPORT jstring JNICALL Java_com_sevenlib_app_imageviewer_MainAct_getMessage(JNIEnv *env, jobject thisObj) {
    __android_log_print(ANDROID_LOG_INFO, "duy", "Call from native side! I'm in HEP");
    return (*env)->NewStringUTF(env, "Hello from native code!");
}
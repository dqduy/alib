/*
 * Author: Duy Quoc
 *
 * Description: Parse characters from file with utf-8 encoding
 *
 * Detail: 
 *      - English
 *      - Vietnamese
 *      - Chinese
 */
 
//#define WINDOWS -1

#if defined(WINDOWS)
#include<windows.h>
#endif
#include<iostream>
#include<cstdint>
#include<vector>
#include<fstream>
#include<cstring>
#include "CodePoint.h"

using namespace std;

vector<CodePoint*>      source;         //Extract code points (utf8 encoding) from file
vector<CodePoint*>      charList;       //List of code point after analyze from source string
vector<int>             charListCount;  //Amount of each code point in charList
char                    sourceName[128];
int                     bytes = 0;

void loadString() 
{
    char c;
    CodePoint* cp = nullptr;
    ifstream inputSource(sourceName, ios::binary);

    while(true) {
        if(inputSource.get(c)) {
            ++bytes;
            if((c & U8Mask::maskClass1e) == 0x0) {
                if(cp != nullptr && cp->size() > 0)
                    source.push_back(cp);
                cp = new CodePoint();
                cp->addCodeUnit(c);
                //cout<< "case 1 - " << (int) text[index] << " - " << (text[0] & maskClass1e) <<endl;
            } else if(((c & U8Mask::maskClass2e) == 0xC0) || 
                      ((c & U8Mask::maskClass3e) == 0xE0) ||
                      ((c & U8Mask::maskClass4e) == 0xF0)) {
                //cout<< "case 2 - " << (int) text[index] <<endl;
                if(cp != nullptr && cp->size() > 0)
                    source.push_back(cp);
                cp = new CodePoint();
                cp->addCodeUnit(c);
            } else if(((c & U8Mask::maskClassee) == 0x80)) {
                cp->addCodeUnit(c);
                //cout<< "case 3 - " << (int) text[index] <<endl;
            }
        }
        else {
            source.push_back(cp);
            break;
        } 
    }
}

void sortString() 
{
    for(int out = 0; out < charList.size() - 1; ++out) {
        for(int in = out + 1; in < charList.size(); ++in) {
            if(charList[out]->getCodePoint() > charList[in]->getCodePoint()) {
                CodePoint* tmp = charList[out];
                charList[out] = charList[in];
                charList[in] = tmp;

                int tmp1 = charListCount[out];
                charListCount[out] = charListCount[in];
                charListCount[in] = tmp1;
            }
        }
    }
}

int hasCodePoint(int cp) 
{
    for(int index = 0; index < charList.size(); ++index) {
        if(cp == charList[index]->getCodePoint())
            return index;
    }
    return -1;
}

void analyzeString() 
{
    int contain = -1;
    for(int out = 0; out < source.size(); ++out) {
        contain = hasCodePoint(source[out]->getCodePoint());
        if(contain != -1)
            charListCount[contain]++;
        else {
            charList.push_back(source[out]);
            charListCount.push_back(1);
        }
    }

    sortString();
}

void displayString()
{
    for(int index = 0; index < charListCount.size(); ++index) {
        cout<< "Code point: U+" << std::hex << charList[index]->getCodePoint()
            << "\t" << std::dec << charList[index]->getCodePoint() 
            << "\t" << charList[index]->getString()
            << "\t\tappear\t " << std::dec << charListCount[index] <<endl;
    }
    
    cout<< "Total read bytes in file: " << bytes << "\n";
    cout<< "Total cp from file : " << source.size() << "\n";
    cout<< "Total cp after sort: " << charList.size() << "\n";
}

int main(int argc, char* argv[])
{
#if defined(WINDOWS)
    SetConsoleOutputCP(65001);
#endif    
    std::strcpy(sourceName, argv[1]);
	loadString();
    analyzeString();
    displayString();

    return 0;
}
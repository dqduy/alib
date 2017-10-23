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

using namespace std;

#include<iostream>
#include<cstdint>
#include<vector>
#include<bitset>
#include<fstream>

int len = 11;
uint8_t* text = new uint8_t[len];
int32_t cp[4];

uint8_t maskClass1e = 0x80;     //1000 0000
uint8_t maskClass2e = 0xE0;     //1110 0000
uint8_t maskClass3e = 0xF0;     //1111 0000
uint8_t maskClass4e = 0xF8;     //1111 1000
uint8_t maskClassee = 0xC0;     //1100 0000

uint8_t maskClass1d = 0x7F;     //0111 1111
uint8_t maskClass2d = 0x1F;     //0001 1111
uint8_t maskClass3d = 0x0F;     //0000 1111
uint8_t maskClass4d = 0x07;     //0000 0111
uint8_t maskClassed = 0x3F;     //0011 1111

struct CodePoint {
    vector<uint8_t> codeUnits;
    void addCodeUnit(uint8_t unit)
    {
        codeUnits.push_back(unit);
    }

    uint32_t getCodePoint() 
    {
        uint32_t cp = 0x0;
        switch(size()) {
            case 1:
                cp = codeUnits[0] & maskClass1d;
                break;
            case 2:
                cp = codeUnits[0] & maskClass2d;
                cp = (cp << 6) | (codeUnits[1] & maskClassed);
                break;
            case 3:
                cp = codeUnits[0] & maskClass3d;
                cp = (cp << 6) | (codeUnits[1] & maskClassed);
                cp = (cp << 6) | (codeUnits[2] & maskClassed);
                break;
            case 4:
                cp = codeUnits[0] & maskClass4d;
                cp = (cp << 6) | (codeUnits[1] & maskClassed);
                cp = (cp << 6) | (codeUnits[2] & maskClassed);
                cp = (cp << 6) | (codeUnits[3] & maskClassed);
                break;
        }
        return cp;
    }

    int size() 
    {
        return codeUnits.size();
    }

    void print() 
    {
        cout<< "This code point: " << getCodePoint() <<endl;
        cout<< "Code units: " << size() <<endl;
        for(int index = 0; index < size(); ++index) {
            cout<< "Unit " << index << ": "
                << bitset<8>(codeUnits[index]) << "\t"
                << (int)codeUnits[index] << "\n";
            //cout<< std::hex << codeUnits[index] << "\n";
        }
        cout<<endl;    
    }
};

vector<CodePoint*> str;

void readString()
{
    //sample text: đặng
    //code point: đ     U+0111  - 273   - 0001 0001 0001
    //  code unit       0xC4    - 196   - 11000100 
    //                  0x91    - 145   - 10010001
    //
    //code point: ặ     U+1EB7  - 7863  - 0001 1110 1011 0111
    //  code unit       0xE1    - 225   - 1110 0001
    //                  0xBA    - 186   - 1011 1010
    //                  0xB7    - 183   - 1011 0111
    //
    //code point: n     U+006E  - 110
    //  code unit       0x6E    - 110   - 0110 1110
    //
    //code point: g     U+0067  - 103
    //  code unit       0x67    - 103   - 0110 0111
    //
    //67872
    //
    text[0] = 0xC4;
    text[1] = 0x91;
    text[2] = 0xE1;
    text[3] = 0xBA;
    text[4] = 0xB7;
    text[5] = 0x6E;
    text[6] = 0x67;
    text[7] = 0xF0;
    text[8] = 0x90;
    text[9] = 0xA4;
    text[10] = 0xA0;
}

void parseString() 
{   
    CodePoint* cp = nullptr;//new CodePoint();

    for(int index = 0; index < len; ++index) 
    {
        if((text[index] & maskClass1e) == 0x0) {
            if(cp != nullptr && cp->size() > 0)
                str.push_back(cp);
            cp = new CodePoint();
            cp->addCodeUnit(text[index]);

            if(index >= len - 1)
                str.push_back(cp);
            //cout<< "case 1 - " << (int) text[index] << " - " << (text[0] & maskClass1e) <<endl;
        } else if(((text[index] & maskClass2e) == 0xC0) || 
                ((text[index] & maskClass3e) == 0xE0) ||
                ((text[index] & maskClass4e) == 0xF0)) {
            //cout<< "case 2 - " << (int) text[index] <<endl;
            if(cp != nullptr && cp->size() > 0)
                str.push_back(cp);
            cp = new CodePoint();
            cp->addCodeUnit(text[index]);
        } else if(((text[index] & maskClassee) == 0x80)) {
            cp->addCodeUnit(text[index]);
            //cout<< "case 3 - " << (int) text[index] <<endl;
            if(index >= len - 1)
                str.push_back(cp);
        }

    }
    //cout<< text <<endl;

}

void printString()
{
    for(int index = 0; index < str.size(); ++index)
        str[index]->print();
    cout<< "Total cp: " << str.size() <<endl;
}

#include<sys/stat.h>
#include<stdio.h>

vector<CodePoint*>      source;         //Extract and format code point (utf8 encoding) from file

vector<CodePoint*>      charList;       //List of code point after analyze from source string
vector<int>             charListCount;  //Amount of each code point in charList
char                   sourceName[26] = "sample_vi.txt";


void loadString() 
{
    bool isEndFile = false;
    char c;
    int count = 0;
    CodePoint* cp = nullptr;
    ifstream inputSource(sourceName);

    while(true) 
    {
        if(inputSource.get(c))
        {
            ++count;
            printf("%d ", (unsigned char)c);
            if((c & maskClass1e) == 0x0) {
                if(cp != nullptr && cp->size() > 0)
                    source.push_back(cp);
                cp = new CodePoint();
                cp->addCodeUnit(c);

//                if(index >= len - 1)
//                    str.push_back(cp);
                //cout<< "case 1 - " << (int) text[index] << " - " << (text[0] & maskClass1e) <<endl;
            } else if(((c & maskClass2e) == 0xC0) || 
                      ((c & maskClass3e) == 0xE0) ||
                      ((c & maskClass4e) == 0xF0)) {
                //cout<< "case 2 - " << (int) text[index] <<endl;
                if(cp != nullptr && cp->size() > 0)
                    source.push_back(cp);
                cp = new CodePoint();
                cp->addCodeUnit(c);
            } else if(((c & maskClassee) == 0x80)) {
                cp->addCodeUnit(c);
                //cout<< "case 3 - " << (int) text[index] <<endl;
//                if(index >= len - 1)
//                    str.push_back(cp);
            }

        }
        else 
        {
            source.push_back(cp);
            cout<< "End of file" <<endl;
            break;
        } 

    }

    cout<< "Total bytes: " << count <<endl;
}

void extractString() 
{
}

void analyzeString() 
{
}

void displayString()
{
    for(int index = 0; index < source.size(); ++index)
        source[index]->print();
    cout<< "Total cp: " << source.size() <<endl;
}

int main()
{   
    //readString();
    //parseString();
    //printString();
    //cout<< sizeof(int);
    //
    loadString();
    displayString();

    struct stat r;
    if(stat("sample_vi.txt", &r) == 0) {
        cout<< "Total size on disk: " << r.st_size <<endl; 
    }

    //    cout << "bytes: " << i <<endl;

    return 0;
}

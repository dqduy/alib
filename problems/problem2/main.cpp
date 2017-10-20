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

struct CodePoint {
    vector<uint8_t> codeUnits;
    void addCodeUnit(uint8_t unit){
        codeUnits.push_back(unit);
    }

    int getCodePoint() {

    }
};

vector<CodePoint*> str;

uint8_t* text = new uint8_t[7];
int32_t cp[4];

uint8_t maskClass1e = 0x00;     //1000 0000
uint8_t maskClass2e = 0xE0;     //1110 0000
uint8_t maskClass3e = 0xF0;     //1111 0000
uint8_t maskClass4e = 0xF8;     //1111 1000
uint8_t maskClassee = 0xC0;     //1100 0000

uint8_t maskClass1d = 0x7F;     //0111 1111
uint8_t maskClass2d = 0x1F;     //0001 1111
uint8_t maskClass3d = 0x0F;     //0000 1111
uint8_t maskClass4d = 0x07;     //0000 0111
uint8_t maskClassed = 0x3F;     //0011 1111

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
    text[0] = 0xC4;
    text[1] = 0x91;
    text[2] = 0xE1;
    text[3] = 0xBA;
    text[4] = 0xB7;
    text[5] = 0x6E;
    text[6] = 0x67;

}

void parseString() 
{
    uint32_t tmp = text[2] & maskClass3d;
    uint32_t tmq = text[3] & maskClassed;
    uint32_t tmr = text[4] & maskClassed;

    //cout<< ((((tmp << 6) | tmq) << 6) | tmr) << " ========" <<endl;

    CodePoint* cp = new CodePoint();

    for(int index = 0; index < 7; ++index) 
    {
        if((text[index] & maskClass1e) == 0x0) {
            cp->addCodeUnit(text[index]);
//            str.push_back(cp);
        } else if(((text[index] & maskClass2e) == 0xC0) || 
                  ((text[index] & maskClass3e) == 0xE0) ||
                  ((text[index] & maskClass4e) == 0xF0)) {
            //str.push_back(cp);
            cp = new CodePoint();
            cp->addCodeUnit(text[index]);
        } else if(((text[index] & maskClassee) == 0x80)) {
            cp->addCodeUnit(text[index]);
        }
        
    }

    cout<< str.size();
}

void printString()
{
    cout<< "This string is:" << endl;
    for(int index = 0; index < 7; ++index)
    {
        cout<< (int)text[index] << ", ";
    }

    cout<< endl;

    cout<< 0x0111 << " - " << 0x1EB7 << "- " << 0x006E << " - " << 0x0067 << endl;
}


int main()
{

    unsigned char mask1 = 224;          //1110 0000
    unsigned char mask2 = 31;           //0001 1111

    unsigned char pattern1 = 223;       //1101 1111      
    unsigned char pattern2 = 143;       //1000 1111

    uint32_t p = pattern1 & mask1; 

    //cout<< (p << 1) << "" << (pattern1 & mask2) << endl;

    readString();
    parseString();
    //printString();

    return 0;
}

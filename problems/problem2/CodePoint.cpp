#include "CodePoint.h"
#include<iostream>
#include<bitset>

CodePoint::CodePoint() {

}

void CodePoint::addCodeUnit(uint8_t unit) {
    codeUnits.push_back(unit);
}

uint32_t CodePoint::getCodePoint() {
    uint32_t cp = 0x0;
    switch(size()) {
        case 1:
            cp = codeUnits[0] & U8Mask::maskClass1d;
            break;
        case 2:
            cp = codeUnits[0] & U8Mask::maskClass2d;
            cp = (cp << 6) | (codeUnits[1] & U8Mask::maskClassed);
            break;
        case 3:
            cp = codeUnits[0] & U8Mask::maskClass3d;
            cp = (cp << 6) | (codeUnits[1] & U8Mask::maskClassed);
            cp = (cp << 6) | (codeUnits[2] & U8Mask::maskClassed);
            break;
        case 4:
            cp = (cp << 6) | (codeUnits[2] & U8Mask::maskClassed);
            cp = codeUnits[0] & U8Mask::maskClass4d;
            cp = (cp << 6) | (codeUnits[1] & U8Mask::maskClassed);
            cp = (cp << 6) | (codeUnits[3] & U8Mask::maskClassed);
            break;
    }
    return cp;
}

int CodePoint::size() {
    return codeUnits.size();
}

uint8_t* CodePoint::getContent() {
    uint8_t* s = new uint8_t[size()];
    for(int index = 0; index < size(); ++index)
        s[index] = codeUnits[index];

    return s;
}

uint8_t* CodePoint::getString() {
    uint8_t* s = new uint8_t[size() + 1];
    for(int index = 0; index < size(); ++index)
        s[index] = codeUnits[index];
    s[size()] = 0;

    return s;
}

void CodePoint::print() {
    std::cout<< "This code point: 0x" << std::hex << getCodePoint() 
        << " - " << std::dec << getCodePoint() <<std::endl;
    std::cout<< "Code units: " << size() <<std::endl;
    for(int index = 0; index < size(); ++index) {
        std::cout<< "Unit " << index << ": "
                 << std::bitset<8>(codeUnits[index]) << "\t"
                 << "0x" << std::hex << (int)codeUnits[index]
                 << "  " << std::dec << (int)codeUnits[index]
                 << "\n";
    }
    std::cout<<std::endl;    
}

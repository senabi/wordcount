import { faker } from '@faker-js/faker';
import * as fs from 'fs';
let res = null

console.time()
// generates ~= 6.5GB 
res = faker.lorem.paragraphs(2000000)
fs.writeFileSync('./lorem.txt', res , { flag: 'a' })
fs.writeFileSync('./lorem.txt', '\n', { flag: 'a' })
console.timeEnd()
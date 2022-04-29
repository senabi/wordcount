import { faker } from '@faker-js/faker';
import * as fs from 'fs';
let res = null

console.time()
fs.writeFileSync('./lorem.txt', "")
// generates ~= 5GB 
let nIters = 2604
for (let i = 0; i < nIters ; i++) {
  res = faker.lorem.paragraphs(10000)
  fs.writeFileSync('./lorem.txt', res , { flag: 'a' })
}
fs.writeFileSync('./lorem.txt', '\n', { flag: 'a' })
console.timeEnd()
import { promises as fs } from 'fs';
import path from 'path';
import markdownIt from 'markdown-it';

const md = new markdownIt();

export default async function handler(req, res) {
    const { slug } = req.query;
    const filePath = path.join(process.cwd(), 'posts', `${slug}.md`);
    
    try {
      const fileContent = await fs.readFile(filePath, 'utf8');
      const htmlContent = md.render(fileContent);
  
      res.setHeader('Content-Type', 'text/html');
      res.status(200).send(htmlContent);
    } catch (err) {
      res.status(404).send('Post not found');
    }
  }
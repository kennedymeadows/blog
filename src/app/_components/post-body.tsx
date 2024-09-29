import markdownStyles from "./markdown-styles.module.css";

type Props = {
  content: string;
  tags: string[] | undefined;
};

export function PostBody({ content, tags }: Props) {
  return (
    <div className="max-w-2xl mx-auto">
      <div
        className={markdownStyles["markdown"]}
        dangerouslySetInnerHTML={{ __html: content }}
      />
        {tags && (
          <div className="mt-8">
            <h3 className="mb-4 pointer-events-none">Tags:</h3>
            <ul className="flex space-x-2">
              {tags.map((tag: string) => (
                <li key={tag} className="bg-gray-200 p-2 rounded pointer-events-none">
                  {tag}
                </li>
              ))}
            </ul>
          </div>
        )}
    </div>
  );
}

import { EnliTaskPage } from './app.po';

describe('enli-task App', () => {
  let page: EnliTaskPage;

  beforeEach(() => {
    page = new EnliTaskPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!');
  });
});

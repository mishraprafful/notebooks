import { pageNotfound } from "~/__tests__/cypress/cypress/pages/pageNotFound";
import { home } from "~/__tests__/cypress/cypress/pages/home";

describe('Application', () => {

    it('Page not found should render', () => {
        pageNotfound.visit()
    });

    it('Home page should have primary button', () => {
        home.visit()
        home.findButton();
    });

});
